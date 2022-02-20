import { ApiProperty } from '@nestjs/swagger';

export class CreateCardDto {
  @ApiProperty({ example: 'Ivan', description: 'FirstName' })
  readonly firstName: string;
  @ApiProperty({ example: 'Ivanov', description: 'LastName' })
  readonly lastName: string;
  @ApiProperty({ example: 1234567891234567, description: 'CardNumber' })
  readonly cardNumber: number;
  @ApiProperty({ example: 'visa', description: 'Company' })
  readonly company: string;
  @ApiProperty({ example: 12, description: 'MM' })
  readonly mm: number;
  @ApiProperty({ example: 22, description: 'YY' })
  readonly yy: number;
  @ApiProperty({ example: '123', description: 'CVV' })
  readonly cvv: string;
}
