"use client";

import { Box } from "@chakra-ui/react";
import { HomePage } from "@/components/templates/HomePage";

export default function Home() {
  return (
    <Box minH="100vh" bg="gray.900">
      <HomePage />
    </Box>
  );
}
