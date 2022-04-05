import { Column, DataType, Model, Table } from 'sequelize-typescript';
import { ApiProperty } from '@nestjs/swagger';

interface CardsCreationAttrs {
  firstName: string;
  lastName: string;
  email: string;
  cardNumber: string;
  company: string;
  mm: string;
  yy: string;
  cvv: string;
}

@Table({ tableName: 'cards' })
export class CardsModel extends Model<CardsModel, CardsCreationAttrs> {
  @ApiProperty({ example: '1', description: 'Unique id' })
  @Column({
    type: DataType.INTEGER,
    unique: true,
    autoIncrement: true,
    primaryKey: true,
  })
  id: number;
  @ApiProperty({ example: 'Ivan', description: 'FirstName' })
  @Column({
    type: DataType.STRING,
    allowNull: false,
  })
  firstName: string;
  @ApiProperty({ example: 'Ivanov', description: 'LastName' })
  @Column({
    type: DataType.STRING,
    allowNull: false,
  })
  lastName: string;
  @ApiProperty({ example: 'ivanov@mail.ru', description: 'Email' })
  @Column({
    type: DataType.STRING,
    allowNull: false,
  })
  email: string;
  @ApiProperty({ example: '1234567891234567', description: 'CardNumber' })
  @Column({
    type: DataType.STRING,
    allowNull: false,
  })
  cardNumber: string;
  @ApiProperty({ example: 'visa', description: 'Company' })
  @Column({
    type: DataType.STRING,
    allowNull: false,
  })
  company: string;
  @ApiProperty({ example: '12', description: 'MM' })
  @Column({
    type: DataType.STRING,
    allowNull: false,
  })
  mm: string;
  @ApiProperty({ example: '2022', description: 'YY' })
  @Column({
    type: DataType.STRING,
    allowNull: false,
  })
  yy: string;
  @ApiProperty({ example: '123', description: 'Cvv' })
  @Column({
    type: DataType.STRING,
    allowNull: false,
  })
  cvv: string;
}
