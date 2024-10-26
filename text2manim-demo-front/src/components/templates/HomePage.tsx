import { useState } from "react";
import { useRouter } from "next/navigation";
import { Container, VStack } from "@chakra-ui/react";
import { PromptSection } from "@/components/organisms/PromptSection";
import { RecentVideosCarousel } from "@/components/organisms/RecentVideosCarousel";
import { EmailModal } from "@/components/molecules/EmailModal";
import { toaster } from "@/components/atoms/chakra/toaster";
import { SettingsDrawer } from "@/components/organisms/SettingDrawer";
export function HomePage() {
  const router = useRouter();
  const [isEmailModalOpen, setIsEmailModalOpen] = useState(false);
  const [currentPrompt, setCurrentPrompt] = useState("");
  const [isLoading, setIsLoading] = useState(false);

  const handlePromptSubmit = async (prompt: string) => {
    if (!prompt.trim()) {
      toaster.create({
        title: "エラー",
        description: "プロンプトを入力してください",
        type: "error",
      });
      return;
    }

    const storedEmail = localStorage.getItem("userEmail");
    if (!storedEmail) {
      setCurrentPrompt(prompt);
      setIsEmailModalOpen(true);
      return;
    }
    await submitGeneration(prompt, storedEmail);
  };

  const handleEmailSubmit = async (email: string) => {
    localStorage.setItem("userEmail", email);
    setIsEmailModalOpen(false);
    await submitGeneration(currentPrompt, email);
  };

  const submitGeneration = async (prompt: string, email: string) => {
    setIsLoading(true);
    try {
      const response = await fetch(
        "https://api.text2manim-demo.kawakin.tech/v1/generations",
        {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify({ prompt, email }),
        },
      );

      // エラーレスポンスの詳細なハンドリング
      if (!response.ok) {
        const errorData = await response.json().catch(() => ({}));
        throw new Error(errorData.message || `エラー: ${response.status}`);
      }

      const data = await response.json();
      if (!data.request_id) {
        throw new Error("リクエストIDが見つかりません");
      }

      router.push(`/generations/${data.request_id}`);
      toaster.create({
        title: "動画生成リクエストを受け付けました",
        description: "生成まで1分ほどかかります",
        type: "success",
      });
      localStorage.removeItem("prompt");
    } catch (error) {
      setIsLoading(false); // NOTE: 失敗した時だけローディングか解除, 成功したらそのまま遷移するから
      toaster.create({
        title: "動画生成リクエストに失敗しました",
        description:
          error instanceof Error
            ? error.message
            : "時間をおいてから再度お試しください",
        type: "error",
      });
    }
  };

  return (
    <Container maxW="container.xl" py={10}>
      <VStack padding={10}>
        <PromptSection isLoading={isLoading} onSubmit={handlePromptSubmit} />
        <RecentVideosCarousel />
      </VStack>
      <EmailModal
        isOpen={isEmailModalOpen}
        onClose={() => setIsEmailModalOpen(false)}
        onSubmit={handleEmailSubmit}
      />
      <SettingsDrawer />
    </Container>
  );
}
