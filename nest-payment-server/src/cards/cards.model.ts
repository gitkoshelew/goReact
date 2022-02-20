import { Column, DataType, Model, Table } from 'sequelize-typescript';
import { ApiProperty } from '@nestjs/swagger';

interface CardsCreationAttrs {
  firstName: string;
  lastName: string;
  cardNumber: number;
  company: string;
  mm: number;
  yy: number;
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
  @ApiProperty({ example: 1234567891234567, description: 'CardNumber' })
  @Column({
    type: DataType.BIGINT,
    allowNull: false,
  })
  cardNumber: number;
  @ApiProperty({ example: 'visa', description: 'Company' })
  @Column({
    type: DataType.STRING,
    allowNull: false,
  })
  company: string;
  @ApiProperty({ example: 12, description: 'MM' })
  @Column({
    type: DataType.INTEGER,
    allowNull: false,
  })
  mm: number;
  @ApiProperty({ example: 22, description: 'YY' })
  @Column({
    type: DataType.INTEGER,
    allowNull: false,
  })
  yy: number;
  @ApiProperty({ example: '123', description: 'Cvv' })
  @Column({
    type: DataType.STRING,
    allowNull: false,
  })
  cvv: string;
}
