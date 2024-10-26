import { Button as ChakraButton, type ButtonProps } from "@chakra-ui/react";

export function Button({ children, ...props }: ButtonProps) {
  return (
    <ChakraButton
      colorScheme="teal"
      size="lg"
      fontWeight="bold"
      _hover={{ transform: "translateY(-2px)" }}
      transition="all 0.2s"
      {...props}
    >
      {children}
    </ChakraButton>
  );
}
