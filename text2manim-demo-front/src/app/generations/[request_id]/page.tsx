"use client";

import { useParams } from "next/navigation";
import { Box } from "@chakra-ui/react";
import { GenerationPage } from "@/components/templates/GenerationPage";

export default function GenerationStatusPage() {
  const params = useParams();
  const requestId = params.request_id as string;

  return (
    <Box minH="100vh" bg="gray.50">
      <GenerationPage requestId={requestId} />
    </Box>
  );
}
