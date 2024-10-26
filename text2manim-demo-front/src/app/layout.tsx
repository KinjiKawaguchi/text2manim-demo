import { Inter } from "next/font/google";
import { Provider } from "@/components/atoms/chakra/provider";
import { Toaster } from "@/components/atoms/chakra/toaster";
const inter = Inter({
  subsets: ["latin"],
  display: "swap",
});

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html className={inter.className} lang="ja" suppressHydrationWarning>
      <head />
      <body>
        <Provider>
          <Toaster />
          {children}
        </Provider>
      </body>
    </html>
  );
}
