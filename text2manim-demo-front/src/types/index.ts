export interface GenerationResponse {
  id: string;
  request_id: string;
  prompt: string;
  status: "pending" | "processing" | "completed" | "unspecified" | "failed";
  video_url?: string;
  email: string;
  created_at: string;
  updated_at: string;
}

export interface RecentVideo {
  request_id: string;
  prompt: string;
  video_url: string;
  created_at: string;
}
