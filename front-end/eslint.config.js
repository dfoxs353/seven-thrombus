import { fileURLToPath } from 'url';
import * as path from 'path';

import { fixupPluginRules } from '@eslint/compat';
import { FlatCompat } from '@eslint/eslintrc';
import eslintJs from '@eslint/js';
import eslintTs from 'typescript-eslint';
import eslintReact from 'eslint-plugin-react';

const project = './tsconfig.json';
const filename = fileURLToPath(import.meta.url);
const dirname = path.dirname(filename);
const compat = new FlatCompat({
  baseDirectory: dirname,
  recommendedConfig: eslintJs.configs.recommended
});

function legacyPlugin(name, alias = name) {
  const plugin = compat.plugins(name)[0]?.plugins?.[alias];

  if (!plugin) {
    throw new Error(`Unable to resolve plugin ${name} and/or alias ${alias}`);
  }

  return fixupPluginRules(plugin);
}

export default eslintTs.config(
  eslintJs.configs.recommended,
  ...eslintTs.configs.recommended,
  ...compat.extends('plugin:import/typescript'),
  {
    languageOptions: {
      parserOptions: {
        project,
        tsconfigRootDir: import.meta.dirname
      }
    },
    settings: {
      'import/resolver': {
        typescript: {
          alwaysTryTypes: true,
          project
        }
      }
    },
    plugins: {
      react: eslintReact,
      import: legacyPlugin('eslint-plugin-import', 'import')
    }
  },
  {
    rules: {
      semi: 2,
      eqeqeq: 2,
      'no-console': 2,
      'react/jsx-first-prop-new-line': [2, 'multiline'],
      'import/order': [2, { 'newlines-between': 'always' }],
      'import/newline-after-import': [2, { count: 2 }]
    }
  }
);
