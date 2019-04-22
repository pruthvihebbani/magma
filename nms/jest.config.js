/**
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * @format
 */

const esModules = ['lodash-es'].join('|');

module.exports = {
  collectCoverageFrom: [
    '**/fbcnms-projects/**/*.js',
    '**/fbcnms-packages/**/*.js',
    '!**/__mocks__/**',
    '!**/__tests__/**',
    '!**/node_modules/**',
  ],
  coverageReporters: ['json'],
  moduleNameMapper: {
    '\\.(jpg|jpeg|png|gif|eot|otf|webp|svg|ttf|woff|woff2|mp4|webm|wav|mp3|m4a|aac|oga)$':
      '<rootDir>/__mocks__/fileMock.js',
    '\\.(css|less)$': 'identity-obj-proxy',
  },
  modulePathIgnorePatterns: [],
  projects: ['<rootDir>'],
  testEnvironment: 'jsdom',
  testPathIgnorePatterns: ['/node_modules/'],
  transformIgnorePatterns: [`/node_modules/(?!${esModules})`],
  transform: {
    '^.+\\.js$': 'babel-jest',
  },
};
