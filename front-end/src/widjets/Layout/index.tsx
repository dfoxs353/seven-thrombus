import { ReactNode } from "react";
import { Footer, Header } from "../../shared/ui";

type TLayoutProps = {
    children: ReactNode
}

export const Layout = ({ children }: TLayoutProps) => {
    return (
        <div className={'min-h-screen w-full'}>
            <Header />
            <main>{children}</main>
            <Footer />
        </div>
    )
}