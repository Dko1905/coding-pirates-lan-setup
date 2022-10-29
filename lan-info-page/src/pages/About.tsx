import { FC } from 'react';

export const About: FC<{}> = () => {
  return (
    <div className="text-black dark:text-white">
      Made in 1 day by <a
        className="text-blue-700 underline" href="https://0chaos.eu/"
      >Daniel Florescu</a> for <a
        className="text-blue-700 underline" href="https://www.facebook.com/codingpirateskoege/"
      >Coding Pirates Køge</a>.
    </div>
  )
};
