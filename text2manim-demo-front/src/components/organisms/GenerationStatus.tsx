import { Spinner } from "@/components/atoms/Spinner";
import { ProgressBar, ProgressRoot } from "@/components/atoms/chakra/progress";
import { VideoPlayer } from "@/components/molecules/VideoPlayer";
import type { GenerationResponse } from "@/types";
import { Box, Heading, Text, VStack } from "@chakra-ui/react";
import { useEffect, useState } from "react";

interface Props {
  generation: GenerationResponse | null;
  error: string | null;
}

const promotions = [
  "数式を美しくアニメーション化しています...",
  "数学的な概念を視覚化しています...",
  "説明をわかりやすく構成しています...",
  "最終的な動画を生成しています...",
];

export function GenerationStatus({ generation, error }: Props) {
  const [promotionIndex, setPromotionIndex] = useState(0);

  useEffect(() => {
    if (generation?.status === "processing") {
      const interval = setInterval(() => {
        setPromotionIndex((prev) => (prev + 1) % promotions.length);
      }, 5000);
      return () => clearInterval(interval);
    }
  }, [generation?.status]);

  if (error) {
    return (
      <Box textAlign="center" p={10}>
        <Text color="red.500">エラーが発生しました: {error}</Text>
      </Box>
    );
  }

  if (!generation) {
    return <Spinner />;
  }

  return (
    <VStack padding={8} align="stretch">
      <Box textAlign="center">
        <Heading size="lg" mb={4}>
          {generation.status === "completed"
            ? "動画の生成が完了しました！"
            : "動画を生成しています"}
        </Heading>
        <Text color="gray.600" mb={4}>
          プロンプト: {generation.prompt}
        </Text>
      </Box>

      {generation.status === "completed" && generation.video_url ? (
        <VideoPlayer url={generation.video_url} />
      ) : (
        <VStack padding={6}>
          <ProgressRoot size="lg">
            <ProgressBar colorScheme="teal" />
          </ProgressRoot>
          <Text fontSize="lg" color="gray.600">
            {promotions[promotionIndex]}
          </Text>
        </VStack>
      )}
    </VStack>
  );
}
