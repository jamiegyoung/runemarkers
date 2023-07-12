import fetch from 'node-fetch';
import { strict as assert } from 'node:assert';

import fs from 'fs';

const getTileData = () =>
  fs
    .readdirSync(`./src/tiles`, {
      encoding: `utf-8`,
      withFileTypes: true,
    })
    .filter((file) => file.isFile() && file.name.endsWith(`.json`))
    .map((file) => ({
      safeURI: file.name.replace(`.json`, ``),
      ...JSON.parse(
        fs.readFileSync(`./src/tiles/${file.name}`, { encoding: `utf-8` }),
      ),
    }))
    .map((entity) => ({
      fullName: `${entity.name} ${
        entity.subcategory ? `(${entity.subcategory})` : ``
      }`.trim(),
      thumbnail: entity.thumbnail,
    }));

let lastRequestTime = Date.now();

const processQueue = async (queue) => {
  const now = Date.now();
  const total = queue.length;
  let current = 0;
  while (queue.length > 0) {
    lastRequestTime = now;
    const { fullName, thumbnail } = queue.shift();
    console.log(`processing ${fullName} (${++current}/${total})`);
    await asyncFetch(thumbnail, fullName);
    await new Promise((resolve) => setTimeout(resolve, 200));
  }
};

const asyncFetch = async (url, fullName) => {
  console.log(`Fetching ${url}...`);
  const res = await fetch(url);
  try {
    assert.equal(res.status, 200);
    console.log(`${fullName}'s thumbnail fetched successfully`);
  } catch (error) {
    console.error(`!! ${fullName}'s thumbnail failed to fetch !!`);
    console.error(error);
    process.exit(1);
  }
};

const main = async () => {
  const fetchQueue = getTileData();
  processQueue(fetchQueue);
  // console.log(fetchQueue);
};

main();
