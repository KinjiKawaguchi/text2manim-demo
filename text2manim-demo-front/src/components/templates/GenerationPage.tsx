import { Container } from "@chakra-ui/react";
import { GenerationStatus } from "@/components/organisms/GenerationStatus";
import { useGenerationStatus } from "@/hooks/useGenerationStatus";

interface Props {
  requestId: string;
}

export function GenerationPage({ requestId }: Props) {
  const { generation, error } = useGenerationStatus(requestId);

  return (
    <Container maxW="container.xl" py={10}>
      <GenerationStatus generation={generation} error={error} />
    </Container>
  );
}
