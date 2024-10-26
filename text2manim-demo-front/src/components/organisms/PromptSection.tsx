import { PromptForm } from "@/components/molecules/PromptForm";
import { Box, Heading, Text, VStack } from "@chakra-ui/react";

interface Props {
  isLoading: boolean;
  onSubmit: (prompt: string) => void;
}

export function PromptSection({ isLoading, onSubmit }: Props) {
  return (
    <Box w="100%" maxW="800px" mx="auto" p={6}>
      <VStack padding={6} align="center" textAlign="center">
        <Heading size="xl">Text2Manim Demo</Heading>
        <Text fontSize="lg" color="gray.600">
          プロンプトを入力して、数学の説明動画を自動生成しましょう
        </Text>
        <PromptForm isLoading={isLoading} onSubmit={onSubmit} />
      </VStack>
    </Box>
  );
}
