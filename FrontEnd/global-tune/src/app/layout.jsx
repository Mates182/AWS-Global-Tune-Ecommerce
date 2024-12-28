import Navbar from "@/components/Navbar";
import Breadcrumb from "@/components/Breadcrumb";
import { Roboto } from "next/font/google";
import "./globals.css";
import "bootstrap/dist/css/bootstrap.min.css";

const roboto = Roboto({
  weight: ["300", "400", "500", "700"],
  styles: ["italic", "normal"],
  subsets: ["latin"],
});

export const metadata = {
  title: "GlobalTune",
  description:
    "An international e-commerce platform for musical instruments (Project only, not a real store)",
};

export default function RootLayout({ children }) {
  return (
    <html lang="en">
      <body className={roboto.className}>
        <header>
          <Navbar></Navbar>
        </header>
        <main className="container" style={{height: '1000px'}}>
          <Breadcrumb />
          {children}
        </main>
        <footer></footer>
      </body>
    </html>
  );
}
