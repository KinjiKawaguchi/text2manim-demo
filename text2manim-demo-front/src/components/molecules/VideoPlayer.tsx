import { Box, AspectRatio } from "@chakra-ui/react";

interface Props {
  url: string;
}

export function VideoPlayer({ url }: Props) {
  return (
    <Box w="100%" maxW="800px" mx="auto" borderRadius="lg" overflow="hidden">
      <AspectRatio ratio={16 / 9}>
        <video
          controls
          src={url}
          style={{ width: "100%", height: "100%", objectFit: "contain" }}
        >
          <track kind="captions" srcLang="ja" label="日本語" />
        </video>
      </AspectRatio>
    </Box>
  );
}
