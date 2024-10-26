import { Input, VStack } from "@chakra-ui/react";
import { useState } from "react";
import { Button } from "@/components/ui/button";
import {
  DialogBody,
  DialogContent,
  DialogFooter,
  DialogHeader,
  DialogRoot,
} from "@/components/ui/dialog";
import { Fieldset } from "@chakra-ui/react";
import { Field } from "@/components/ui/field";

interface Props {
  isOpen: boolean;
  onClose: () => void;
  onSubmit: (email: string) => void;
}

export function EmailModal({ isOpen, onClose, onSubmit }: Props) {
  const [email, setEmail] = useState("");
  const [isLoading, setIsLoading] = useState(false);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!email.trim()) return;

    setIsLoading(true);
    try {
      await onSubmit(email);
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <DialogRoot open={isOpen} onOpenChange={onClose}>
      <DialogContent as="form" onSubmit={handleSubmit}>
        <DialogHeader>メールアドレスの登録</DialogHeader>
        <DialogBody>
          <VStack padding={4}>
            <Fieldset.Root>
              <Field label="メールアドレス">
                <Input
                  type="email"
                  placeholder="your@email.com"
                  value={email}
                  onChange={(e) => setEmail(e.target.value)}
                />
              </Field>
            </Fieldset.Root>
          </VStack>
        </DialogBody>
        <DialogFooter>
          <Button variant="ghost" mr={3} onClick={onClose}>
            キャンセル
          </Button>
          <Button
            type="submit"
            colorScheme="teal"
            loading={isLoading}
            disabled={!email.trim()}
          >
            登録して生成
          </Button>
        </DialogFooter>
      </DialogContent>
    </DialogRoot>
  );
}
