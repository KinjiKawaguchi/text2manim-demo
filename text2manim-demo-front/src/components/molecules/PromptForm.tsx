import { Textarea } from "@/components/atoms/Textarea";
import { Button } from "@/components/atoms/chakra/button";
import { VStack } from "@chakra-ui/react";
import { useEffect, useState } from "react";

interface Props {
  isLoading: boolean;
  onSubmit: (prompt: string) => void;
}

export function PromptForm({ isLoading, onSubmit }: Props) {
  const [prompt, setPrompt] = useState("");

  const handleChangePrompt = (e: React.ChangeEvent<HTMLTextAreaElement>) => {
    localStorage.setItem("prompt", e.target.value);
    setPrompt(e.target.value);
  };

  useEffect(() => {
    const savedPrompt = localStorage.getItem("prompt");
    if (savedPrompt) {
      setPrompt(savedPrompt);
    }
  }, []);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!prompt.trim()) return;

    onSubmit(prompt);
  };

  return (
    <form onSubmit={handleSubmit} style={{ width: "100%" }}>
      <VStack padding={4} align="stretch">
        <Textarea
          value={prompt}
          onChange={handleChangePrompt}
          disabled={isLoading}
        />
        <Button
          loading={isLoading}
          type="submit"
          loadingText="生成中..."
          disabled={!prompt.trim()}
        >
          動画を生成
        </Button>
      </VStack>
    </form>
  );
}
