import { useState } from "react";
import { useRouter } from "next/navigation";
import { Container, VStack } from "@chakra-ui/react";
import { PromptSection } from "@/components/organisms/PromptSection";
import { RecentVideosCarousel } from "@/components/organisms/RecentVideosCarousel";
import { EmailModal } from "@/components/molecules/EmailModal";

export function HomePage() {
  const router = useRouter();
  const [isEmailModalOpen, setIsEmailModalOpen] = useState(false);
  const [currentPrompt, setCurrentPrompt] = useState("");

  const handlePromptSubmit = async (prompt: string) => {
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
    try {
      const response = await fetch(
        "https://api.text2manim-demo.kawakin.tech/v1/generations",
        {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify({ prompt, email }),
        },
      );
      const data = await response.json();
      router.push(`/generations/${data.request_id}`);
    } catch (error) {
      console.error("Generation request failed:", error);
    }
  };

  return (
    <Container maxW="container.xl" py={10}>
      <VStack padding={10}>
        <PromptSection onSubmit={handlePromptSubmit} />
        <RecentVideosCarousel />
      </VStack>
      <EmailModal
        isOpen={isEmailModalOpen}
        onClose={() => setIsEmailModalOpen(false)}
        onSubmit={handleEmailSubmit}
      />
    </Container>
  );
}
