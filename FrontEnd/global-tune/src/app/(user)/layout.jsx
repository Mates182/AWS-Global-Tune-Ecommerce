import Navbar from "@/components/Navbar";

export default function RootLayout({ children }) {
  return (
    <>
      <header>
        <Navbar></Navbar>
      </header>
      {children}
    </>
  );
}
