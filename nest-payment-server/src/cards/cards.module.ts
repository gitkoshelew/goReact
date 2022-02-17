import { Module } from '@nestjs/common';
import { CardsService } from './cards.service';
import { CardsController } from './cards.controller';
import { SequelizeModule } from '@nestjs/sequelize';
import { CardsModel } from './cards.model';

@Module({
  providers: [CardsService],
  controllers: [CardsController],
  imports: [SequelizeModule.forFeature([CardsModel])],
})
export class CardsModule {}
