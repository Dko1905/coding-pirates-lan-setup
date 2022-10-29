import { FC, useEffect, useState } from 'react';
import { Outlet, Link } from 'react-router-dom'
import { SunIcon, MoonIcon } from '@heroicons/react/24/outline'
import reactLogo from './assets/react.svg';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';

const App: FC<{}> = () => {
  const queryClient = new QueryClient()
  const [isDarkMode, setIsDarkMode] = useState(false);
  const toggleDarkMode = () => {
    const currentMode = !isDarkMode
    console.log('IsDarkMode: ' + currentMode)
    setIsDarkMode(currentMode);
    localStorage.theme = currentMode ? 'dark' : 'light'
  };
  useEffect(() => {
    if (localStorage.theme === 'dark' || (!('theme' in localStorage) && window.matchMedia('(prefers-color-scheme: dark)').matches)) {
      setIsDarkMode(true)
    } else {
      setIsDarkMode(false);
    }
  });

  return (
    <QueryClientProvider client={queryClient}>
      <div className={(isDarkMode ? 'dark' : '') + ' h-screen antialiased'}>
        <div className="bg-white dark:bg-gdark h-screen transition-colors duration-800 ">
          <header className="bg-blue-400 dark:bg-blue-800 shadow-md sticky top-0 z-50 px-4 py-1">
            <nav className="flex justify-between items-center">
              <Link to="/"><span className="text-2xl text-white">LAN Party UI</span></Link>
              <div className="text-white flex md:justify-end">
                <Link to="/" className="mr-3">Home</Link>
                <Link to="about" className="mr-3">About</Link>
                <button onClick={toggleDarkMode}>
                  {isDarkMode === true
                    ? <SunIcon className="h-6 w-6 text-white"></SunIcon>
                    : <MoonIcon className="h-6 w-6 text-white"></MoonIcon>}
                </button>
              </div>
            </nav>
          </header>
          <div className="mx-2 my-4">
            <Outlet />
          </div>
        </div>
      </div>
    </QueryClientProvider>
  );
};

export default App;
