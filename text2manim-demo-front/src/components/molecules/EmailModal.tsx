import { Input, VStack } from "@chakra-ui/react";
import { useState } from "react";
import { Button } from "@/components/atoms/chakra/button";
import {
  DialogBody,
  DialogContent,
  DialogFooter,
  DialogHeader,
  DialogRoot,
} from "@/components/atoms/chakra/dialog";
import { Fieldset } from "@chakra-ui/react";
import { Field } from "@/components/atoms/chakra/field";

interface Props {
  isOpen: boolean;
  onClose: () => void;
  onSubmit: (email: string) => void;
}

export function EmailModal({ isOpen, onClose, onSubmit }: Props) {
  const [email, setEmail] = useState("");

  // TODO: Emailのバリデーション

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!email.trim()) return;

    onSubmit(email);
  };

  return (
    <DialogRoot open={isOpen} onOpenChange={onClose}>
      <DialogContent>
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
            disabled={!email.trim()}
            //TODO: ボタンにonClickを直接書くのはあまりよくない
            onClick={handleSubmit}
          >
            登録して生成
          </Button>
        </DialogFooter>
      </DialogContent>
    </DialogRoot>
  );
}
