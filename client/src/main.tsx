import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { Provider } from '@/components/ui/provider'
import { Provider as ReduxProvider } from 'react-redux'
import './index.css'
import App from './App.tsx'
import { BrowserRouter } from 'react-router-dom'
import { store } from './store/store.ts'
import { Toaster } from './components/ui/toaster.tsx'

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <Provider>
      <BrowserRouter>
        <ReduxProvider store={store}>
          <Toaster></Toaster>
          <App />
        </ReduxProvider>
      </BrowserRouter>
    </Provider>
  </StrictMode>,
)
