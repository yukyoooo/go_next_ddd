import nextJest from 'next/jest'
const createJestConfig = nextJest({ dir: './' })
const customJestConfig = {
  testPathIgnorePatterns: ['<rootDir>/.next/', '<rootDir>/node_modules/'],
  setupFilesAfterEnv: ['<rootDir>/jest.setup.js'],
  moduleDirectories: ['node_modules', '<rootDir>/src'],
  testEnvironment: 'jsdom',
}
export default createJestConfig(customJestConfig)
