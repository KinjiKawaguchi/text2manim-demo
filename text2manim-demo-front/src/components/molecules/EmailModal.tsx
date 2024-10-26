import { Button } from "@/components/atoms/chakra/button";
import {
  DialogBody,
  DialogContent,
  DialogFooter,
  DialogHeader,
  DialogRoot,
} from "@/components/atoms/chakra/dialog";
import { Field } from "@/components/atoms/chakra/field";
import { toaster } from "@/components/atoms/chakra/toaster";
import { Input, VStack } from "@chakra-ui/react";
import { Fieldset } from "@chakra-ui/react";
import { useState } from "react";

interface Props {
  isOpen: boolean;
  onClose: () => void;
  onSubmit: (email: string) => void;
}

export function EmailModal({ isOpen, onClose, onSubmit }: Props) {
  const [email, setEmail] = useState("");

  const validateEmail = (email: string): boolean => {
    const emailRegex = /^[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,}$/i;

    // 基本的なバリデーション
    if (!email || !email.trim()) {
      toaster.create({
        title: "エラー",
        description: "メールアドレスを入力してください",
        type: "warning",
      });
      return false;
    }

    // メールアドレスの形式チェック
    if (!emailRegex.test(email)) {
      toaster.create({
        title: "エラー",
        description: "正しいメールアドレスの形式で入力してください",
        type: "warning",
      });
      return false;
    }

    // 特定のドメインの制限を追加する場合（オプション）
    const domain = email.split("@")[1];
    const blockedDomains = ["example.com", "test.com"];
    if (blockedDomains.includes(domain)) {
      toaster.create({
        title: "エラー",
        description: "このドメインのメールアドレスは使用できません",
        type: "warning",
      });
      return false;
    }

    return true;
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!validateEmail(email) || !email.trim()) return;

    onSubmit(email);
  };

  return (
    <DialogRoot open={isOpen} onOpenChange={onClose}>
      <DialogContent>
        <DialogHeader>メールアドレスの登録</DialogHeader>
        <DialogBody>
          <VStack padding={4}>
            <Fieldset.Root>
              <Field label="メールアドレス" typeof="email" required>
                <Input
                  type="email"
                  placeholder="your@email.com"
                  value={email}
                  onChange={(e) => setEmail(e.target.value)}
                  required
                />
              </Field>
              <DialogFooter>
                <Button variant="ghost" mr={3} onClick={onClose}>
                  キャンセル
                </Button>
                <Button
                  type="submit"
                  colorScheme="teal"
                  disabled={!email.trim()}
                  onClick={handleSubmit}
                >
                  登録して生成
                </Button>
              </DialogFooter>
            </Fieldset.Root>
          </VStack>
        </DialogBody>
      </DialogContent>
    </DialogRoot>
  );
}
