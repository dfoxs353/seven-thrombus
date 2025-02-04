import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import '../assets/index.css'
import {HomePage} from "./pages/home";
import { Layout } from '../widjets';

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <Layout>
        <HomePage />
    </Layout>
  </StrictMode>,
)
