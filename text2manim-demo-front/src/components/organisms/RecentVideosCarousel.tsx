import { useEffect, useState } from "react";
import {
  Box,
  Heading,
  HStack,
  Text,
  VStack,
  IconButton,
} from "@chakra-ui/react";
import { IoChevronBackOutline, IoChevronForwardOutline } from "react-icons/io5";
import type { RecentVideo } from "@/types";

// モックデータ
const MOCK_VIDEOS: RecentVideo[] = [
  {
    id: "1",
    prompt: "二次方程式の解の公式の導出",
    video_url: "https://example.com/video1.mp4",
    created_at: "2024-10-24T08:00:00Z",
  },
  {
    id: "2",
    prompt: "三角関数の基本的な性質",
    video_url: "https://example.com/video2.mp4",
    created_at: "2024-10-24T07:30:00Z",
  },
  {
    id: "3",
    prompt: "微分の基本的な考え方",
    video_url: "https://example.com/video3.mp4",
    created_at: "2024-10-24T07:00:00Z",
  },
];

export function RecentVideosCarousel() {
  const [videos] = useState<RecentVideo[]>(MOCK_VIDEOS);
  const [currentIndex, setCurrentIndex] = useState(0);

  const handlePrev = () => {
    setCurrentIndex((prev) => (prev - 1 + videos.length) % videos.length);
  };

  const handleNext = () => {
    setCurrentIndex((prev) => (prev + 1) % videos.length);
  };

  useEffect(() => {
    const interval = setInterval(() => {
      setCurrentIndex((prev) => (prev + 1) % videos.length);
    }, 5000);
    return () => clearInterval(interval);
  }, [videos]);

  return (
    <Box w="100%" position="relative" py={8}>
      <Heading size="md" mb={6} textAlign="center">
        最近生成された動画
      </Heading>
      <HStack padding={4} justify="center" align="center">
        <IconButton
          aria-label="Previous video"
          onClick={handlePrev}
          variant="ghost"
        >
          <IoChevronBackOutline />
        </IconButton>
        <Box
          w="300px"
          h="200px"
          position="relative"
          overflow="hidden"
          borderRadius="lg"
          bg="gray.100"
        >
          <VStack
            padding={2}
            p={4}
            position="absolute"
            bottom={0}
            left={0}
            right={0}
            bg="rgba(0, 0, 0, 0.7)"
            color="white"
          >
            <Text fontSize="sm">{videos[currentIndex].prompt}</Text>
          </VStack>
        </Box>
        <IconButton
          aria-label="Next video"
          onClick={handleNext}
          variant="ghost"
        >
          <IoChevronForwardOutline />
        </IconButton>
      </HStack>
    </Box>
  );
}
