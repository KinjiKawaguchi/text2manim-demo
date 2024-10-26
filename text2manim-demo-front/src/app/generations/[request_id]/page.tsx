"use client";

import { GenerationPage } from "@/components/templates/GenerationPage";
import { Box } from "@chakra-ui/react";
import { useParams } from "next/navigation";

export default function GenerationStatusPage() {
  const params = useParams();
  const requestId = params.request_id as string;

  return (
    <Box minH="100vh" bg="gray.900">
      <GenerationPage requestId={requestId} />
    </Box>
  );
}
