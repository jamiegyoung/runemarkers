const { readFileSync, writeFileSync } = require('fs');

const rawData = readFileSync('src/tiles.json', 'utf8');
const parsedData = JSON.parse(rawData);

const sortedData = parsedData.sort((a, b) => {
  if (a.name < b.name) {
    return -1;
  }
  if (a.name > b.name) {
    return 1;
  }
  return 0;
});

writeFileSync('src/tiles.json', JSON.stringify(sortedData, null, 2));
