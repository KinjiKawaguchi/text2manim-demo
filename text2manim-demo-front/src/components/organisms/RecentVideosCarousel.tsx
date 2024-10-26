import type { RecentVideo } from "@/types";
import {
  Box,
  HStack,
  Heading,
  IconButton,
  Text,
  VStack,
  Link,
} from "@chakra-ui/react";
import { useCallback, useEffect, useRef, useState } from "react";
import { IoChevronBackOutline, IoChevronForwardOutline } from "react-icons/io5";
import NextLink from "next/link";

// モックデータを更新してrequest_idを追加
const MOCK_VIDEOS: (RecentVideo & { request_id: string })[] = [
  {
    request_id: "19c11a09-e9e4-4ff5-8f79-f23803eb48fb",
    prompt:
      "円の方程式がなぜ円たらしめるのかをグラフを用いて視覚的に説明する動画を作成してください。",
    video_url:
      "https://storage.googleapis.com/text2manim/videos/19c11a09-e9e4-4ff5-8f79-f23803eb48fb.mp4",
    created_at: "2024-10-24T08:00:00Z",
  },
  {
    request_id: "5bc9acac-4564-469e-aced-e330d1745304",
    prompt: "比例と反比例について説明する動画を作成してください。",
    video_url:
      "https://storage.googleapis.com/text2manim/videos/5bc9acac-4564-469e-aced-e330d1745304.mp4",
    created_at: "2024-10-24T07:30:00Z",
  },
  {
    request_id: "eee286c4-d5db-4eff-883c-8a0a99aa1b7f",
    prompt: "楕円を式とグラフを用いて説明する動画",
    video_url:
      "https://storage.googleapis.com/text2manim/videos/eee286c4-d5db-4eff-883c-8a0a99aa1b7f.mp4",
    created_at: "2024-10-24T07:00:00Z",
  },
];

export function RecentVideosCarousel() {
  const [videos] = useState(MOCK_VIDEOS);
  const [currentIndex, setCurrentIndex] = useState(0);
  const [isPlaying, setIsPlaying] = useState(false);
  const videoRef = useRef<HTMLVideoElement>(null);

  const handlePrev = () => {
    setCurrentIndex((prev) => (prev - 1 + videos.length) % videos.length);
  };

  const handleNext = useCallback(() => {
    setCurrentIndex((prev) => (prev + 1) % videos.length);
  }, [videos.length]);

  // 動画の自動再生と停止を管理
  useEffect(() => {
    const video = videoRef.current;
    if (!video) return;

    const playVideo = async () => {
      try {
        video.currentTime = 0;
        video.muted = true;
        await video.play();
        setIsPlaying(true);
      } catch (error) {
        console.error("動画の再生に失敗しました:", error);
      }
    };

    playVideo();

    const stopTimeout = setTimeout(() => {
      video.pause();
      setIsPlaying(false);
    }, 3000);

    return () => {
      clearTimeout(stopTimeout);
      video.pause();
      setIsPlaying(false);
    };
  }, []);

  useEffect(() => {
    const video = videoRef.current;
    if (!video) return;

    video.currentTime = 0;
    video
      .play()
      .catch((error) => console.error("動画の再生に失敗しました:", error));
  }, []);

  useEffect(() => {
    const interval = setInterval(() => {
      if (!isPlaying) {
        handleNext();
      }
    }, 5000);
    return () => clearInterval(interval);
  }, [isPlaying, handleNext]);

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
          <IoChevronBackOutline size={20} />
        </IconButton>
        <Link
          as={NextLink}
          href={`/generations/${videos[currentIndex].request_id}`}
          _hover={{ textDecoration: "none" }}
        >
          <Box
            w="300px"
            h="200px"
            position="relative"
            overflow="hidden"
            borderRadius="lg"
            bg="gray.100"
            transition="transform 0.2s"
            _hover={{ transform: "scale(1.02)" }}
          >
            <video
              ref={videoRef}
              src={videos[currentIndex].video_url}
              style={{
                width: "100%",
                height: "100%",
                objectFit: "cover",
              }}
              playsInline
            >
              <track kind="captions" srcLang="ja" label="日本語" />
            </video>
            <VStack
              padding={2}
              position="absolute"
              bottom={0}
              left={0}
              right={0}
              bg="rgba(0, 0, 0, 0.7)"
              color="white"
              transition="opacity 0.2s"
              opacity={isPlaying ? 0 : 1}
            >
              <Text fontSize="sm">{videos[currentIndex].prompt}</Text>
            </VStack>
          </Box>
        </Link>
        <IconButton
          aria-label="Next video"
          onClick={handleNext}
          variant="ghost"
        >
          <IoChevronForwardOutline size={20} />
        </IconButton>
      </HStack>
    </Box>
  );
}
