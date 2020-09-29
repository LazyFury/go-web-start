export const randomColor = () => {
  let colors: string[] = [
    '#f50',
    '#2db7f5',
    '#87d068',
    'green',
    'blue',
    'gold',
    'volcano',
    'magenta',
    'red',
    'orange',
    'lime',
    'cyan',
    'geekblue',
    'purple',
  ];
  let random = Math.floor(Math.random() * colors.length);
  let color = colors[random];
  return color;
};

export const emptyPromise = () => new Promise((resolve, reject) => resolve());
