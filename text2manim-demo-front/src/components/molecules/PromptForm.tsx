import { useState } from "react";
import { VStack } from "@chakra-ui/react";
import { Textarea } from "@/components/atoms/Textarea";
import { Button } from "@/components/atoms/chakra/button";

interface Props {
  onSubmit: (prompt: string) => void;
}

export function PromptForm({ onSubmit }: Props) {
  const [prompt, setPrompt] = useState("");
  const [isLoading, setIsLoading] = useState(false);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!prompt.trim()) return;

    setIsLoading(true);
    try {
      await onSubmit(prompt);
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <form onSubmit={handleSubmit} style={{ width: "100%" }}>
      <VStack padding={4} align="stretch">
        <Textarea
          value={prompt}
          onChange={(e) => setPrompt(e.target.value)}
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
