import { Button } from "@/components/atoms/chakra/button";
import {
  DrawerActionTrigger,
  DrawerBackdrop,
  DrawerBody,
  DrawerCloseTrigger,
  DrawerContent,
  DrawerFooter,
  DrawerHeader,
  DrawerRoot,
  DrawerTitle,
  DrawerTrigger,
} from "@/components/atoms/chakra/drawer";
import {
  Box,
  type DrawerOpenChangeDetails,
  Input,
  Text,
  VStack,
} from "@chakra-ui/react";
import { useEffect, useState } from "react";
// src/components/molecules/SettingsDrawer.tsx
import { IoSettings } from "react-icons/io5";

interface SettingsDrawerProps {
  position?: {
    bottom?: number;
    right?: number;
  };
}

export function SettingsDrawer({
  position = { bottom: 4, right: 4 },
}: SettingsDrawerProps) {
  const [isOpen, setIsOpen] = useState(false);
  const [localEmail, setLocalEmail] = useState("");
  const [localPrompt, setLocalPrompt] = useState("");

  useEffect(() => {
    const email = localStorage.getItem("userEmail");
    const prompt = localStorage.getItem("prompt");
    if (email) {
      setLocalEmail(email);
    }
    if (prompt) {
      setLocalPrompt(prompt);
    }
  }, []);

  const handleEmailUpdate = (newEmail: string) => {
    setLocalEmail(newEmail);
    localStorage.setItem("userEmail", newEmail);
  };

  const handlePromptUpdate = (newPrompt: string) => {
    setLocalPrompt(newPrompt);
    localStorage.setItem("prompt", newPrompt);
  };

  const handleClear = () => {
    localStorage.clear();
    setLocalEmail("");
    setLocalPrompt("");
  };

  return (
    <Box position="fixed" bottom={position.bottom} right={position.right}>
      <DrawerRoot
        open={isOpen}
        onOpenChange={(value: DrawerOpenChangeDetails) => setIsOpen(value.open)}
      >
        <DrawerTrigger asChild>
          <Button variant="outline" size="sm">
            <IoSettings size={20} />
          </Button>
        </DrawerTrigger>
        <DrawerContent>
          <DrawerHeader>
            <DrawerTitle>設定</DrawerTitle>
          </DrawerHeader>
          <DrawerBody>
            <VStack padding={4} align="stretch">
              <Box>
                <Text mb={2}>メールアドレス</Text>
                <Input
                  value={localEmail}
                  onChange={(e) => handleEmailUpdate(e.target.value)}
                  placeholder="メールアドレス"
                />
              </Box>
              <Box>
                <Text mb={2}>最後のプロンプト</Text>
                <Input
                  value={localPrompt}
                  onChange={(e) => handlePromptUpdate(e.target.value)}
                  placeholder="プロンプト"
                />
              </Box>
            </VStack>
          </DrawerBody>
          <DrawerFooter>
            <DrawerActionTrigger asChild>
              <Button variant="outline" onClick={handleClear}>
                クリア
              </Button>
            </DrawerActionTrigger>
            <DrawerCloseTrigger asChild>
              <Button>閉じる</Button>
            </DrawerCloseTrigger>
          </DrawerFooter>
        </DrawerContent>
        <DrawerBackdrop />
      </DrawerRoot>
    </Box>
  );
}
