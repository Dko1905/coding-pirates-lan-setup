import CsgoIcon from '../assets/csgo.svg';
import MinecraftIcon from '../assets/minecraft.svg';

export const icons = {
  NONE: 'none',
  CSGO: 'csgo',
  MINECRAFT: 'minecraft'
};

export const nameToIcon = (name: string) => {
  switch (name) {
    case icons.NONE:
      return <span/>
    case icons.CSGO:
      return <CsgoIcon />
      break;
    case icons.MINECRAFT:
      return <MinecraftIcon />
      break;
    default:
      throw Error("Cannot find icon for " + name);
      break;
  }
};
