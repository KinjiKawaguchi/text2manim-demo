"use client";

import { HomePage } from "@/components/templates/HomePage";
import { Box } from "@chakra-ui/react";

export default function Home() {
  return (
    <Box minH="100vh" bg="gray.900">
      <HomePage />
    </Box>
  );
}
