import * as bcrypt from 'bcrypt';

export function encodeCvv(rawCvv: string) {
  const salt = bcrypt.genSaltSync();
  return bcrypt.hashSync(rawCvv, salt);
}
