import './globals.css';

export default function RootLayout({ children }: LayoutProps<'/'>) {
    return (
        <html lang="en">
            <head>
                <link rel="preconnect" href="https://fonts.bunny.net" />
                <link href="https://fonts.bunny.net/css?family=instrument-sans:400,500,600" rel="stylesheet" />
            </head>
            <body>{children}</body>
        </html>
    );
}
