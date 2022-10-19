import { useState, useEffect } from 'react';

function getSessionStorageOrDefault(key: string, defaultValue: string | null) {
  console.log('getSessionStorageOrDefault', { key, defaultValue })
  const stored = window.sessionStorage.getItem(key);
  if (!stored) {
    return defaultValue;
  }
  return JSON.parse(stored);
}

export default function useSessionStorage(key: string, defaultValue: string | null) {
  console.log('useSessionStorage', { key, defaultValue })
  const [value, setValue] = useState(
    getSessionStorageOrDefault(key, defaultValue)
  );

  useEffect(() => {
    window.sessionStorage.setItem(key, JSON.stringify(value));
  }, [key, value]);

  return [value, setValue];
}
