import { ApiProperty } from '@nestjs/swagger';
import { IsAlpha, IsEmail, IsString, Length } from 'class-validator';

export class CreateCardDto {
  @ApiProperty({ example: 'Ivan', description: 'FirstName' })
  @IsString({ message: 'Should be a string' })
  @Length(2, 30, { message: 'Must contain exactly 4 characters' })
  @IsAlpha()
  readonly firstName: string;
  @ApiProperty({ example: 'Ivanov', description: 'LastName' })
  @IsString({ message: 'Should be a string' })
  @Length(2, 30, { message: 'Must contain exactly 4 characters' })
  @IsAlpha()
  readonly lastName: string;
  @ApiProperty({ example: 'ivanov@mail.ru', description: 'Email' })
  @IsString({ message: 'Should be a string' })
  @IsEmail({}, { message: 'Invalid email' })
  readonly email: string;
  @ApiProperty({ example: '1234567891234567', description: 'CardNumber' })
  @IsString({ message: 'Should be a string' })
  @Length(16, 16, { message: 'Must contain exactly 16 characters' })
  readonly cardNumber: string;
  @ApiProperty({ example: 'visa', description: 'Company' })
  @IsString({ message: 'Should be a string' })
  readonly company: string;
  @ApiProperty({ example: '12', description: 'MM' })
  @IsString({ message: 'Should be a number' })
  @Length(2, 2, { message: 'Must contain exactly 2 characters' })
  readonly mm: string;
  @ApiProperty({ example: '2022', description: 'YY' })
  @IsString({ message: 'Should be a number' })
  @Length(4, 4, { message: 'Must contain exactly 4 characters' })
  readonly yy: string;
  @ApiProperty({ example: '123', description: 'CVV' })
  @IsString({ message: 'Should be a string' })
  @Length(3, 3, { message: 'Must contain exactly 3 characters' })
  readonly cvv: string;
}
