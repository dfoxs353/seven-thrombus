import { FieldErrors, FieldValues } from 'react-hook-form';

export const getSuccessMessage = <T extends FieldValues>(
  fieldName: keyof T,
  message: string,
  touchedFields: Partial<Readonly<{ [key in keyof T]: boolean }>>,
  errors: FieldErrors<T>,
  value: unknown
): string | undefined => {
  return touchedFields[fieldName] && !errors[fieldName] && value 
    ? message 
    : undefined;
};