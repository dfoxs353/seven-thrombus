import eslint from '@eslint/js';
import tseslint from '@typescript-eslint/eslint-plugin';
import reactRecommended from 'eslint-plugin-react/configs/recommended.js';
import reactHooks from 'eslint-plugin-react-hooks';
import jsxA11y from 'eslint-plugin-jsx-a11y';
import tanstackQuery from '@tanstack/eslint-plugin-query';
import prettier from 'eslint-plugin-prettier/recommended';

export default tseslint.config(
  {
    ignores: ['**/dist', '**/node_modules'],
  },
  {
    files: ['**/*.{ts,tsx}'],
    ...eslint.configs.recommended,
  },
  {
    languageOptions: {
      parser: tseslint.parser,
      parserOptions: {
        ecmaVersion: 'latest',
        sourceType: 'module',
        project: './tsconfig.json',
        ecmaFeatures: {
          jsx: true,
        },
      },
    },
  },
  {
    plugins: {
      '@typescript-eslint': tseslint,
      'react-hooks': reactHooks,
      'jsx-a11y': jsxA11y,
      '@tanstack/query': tanstackQuery,
    },
    rules: {
      ...tseslint.configs.recommended.rules,
      'no-console': 'warn',
      '@typescript-eslint/no-unused-vars': 'error',
      '@typescript-eslint/no-explicit-any': 'warn',
      '@typescript-eslint/consistent-type-imports': 'error',
      'react-hooks/rules-of-hooks': 'error',
      'react-hooks/exhaustive-deps': 'warn',
      '@tanstack/query/exhaustive-deps': 'error',
      '@tanstack/query/prefer-query-object-syntax': 'error',
    },
  },
  reactRecommended,
  {
    rules: {
      'react/react-in-jsx-scope': 'off',
      'react/jsx-uses-react': 'off',
      'react/prop-types': 'off',
      'jsx-a11y/anchor-is-valid': [
        'error',
        {
          components: ['Link'],
          specialLink: ['to'],
        },
      ],
    },
  },
  prettier,
  {
    files: ['**/*.test.{ts,tsx}'],
    env: {
      jest: true,
    },
  },
);
