import type { GenerationResponse } from "@/types/index";
import { useEffect, useState } from "react";

export function useGenerationStatus(requestId: string) {
  const [generation, setGeneration] = useState<GenerationResponse | null>(null);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const checkStatus = async () => {
      try {
        const response = await fetch(
          `https://api.text2manim-demo.kawakin.tech/v1/generations/${requestId}`,
        );
        const data = await response.json();
        setGeneration(data);

        if (data.status === "pending" || data.status === "processing") {
          setTimeout(checkStatus, 5000);
        }
      } catch (err) {
        setError("Failed to fetch generation status");
      }
    };

    checkStatus();
  }, [requestId]);

  return { generation, error };
}
