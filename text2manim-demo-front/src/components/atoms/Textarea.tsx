import {
  Textarea as ChakraTextarea,
  type TextareaProps,
} from "@chakra-ui/react";

export function Textarea(props: TextareaProps) {
  return (
    <ChakraTextarea
      placeholder="動画を生成するためのプロンプトを入力してください..."
      size="lg"
      minH="150px"
      resize="vertical"
      borderRadius="lg"
      _focus={{
        borderColor: "teal.500",
        boxShadow: "0 0 0 1px var(--chakra-colors-teal-500)",
      }}
      {...props}
    />
  );
}
