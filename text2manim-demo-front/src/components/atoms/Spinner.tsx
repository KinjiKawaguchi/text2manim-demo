import { Spinner as ChakraSpinner, Center } from "@chakra-ui/react";

export function Spinner() {
  return (
    <Center p={10}>
      <ChakraSpinner
        size="xl"
        display="inline-block"
        borderWidth="2px"
        borderStyle="solid"
        borderRadius="full"
        animation="spin"
        animationDuration="slowest"
        borderColor="teal.500"
        borderBottomColor="transparent"
        borderInlineStartColor="transparent"
      />
    </Center>
  );
}
