import { Module } from '@nestjs/common';
import { SequelizeModule } from '@nestjs/sequelize';
import { CardsModule } from './cards/cards.module';
import { ConfigModule } from '@nestjs/config';
import { CardsModel } from './cards/cards.model';

@Module({
  imports: [
    ConfigModule.forRoot(),
    SequelizeModule.forRoot({
      dialect: 'postgres',
      host: process.env.POSTGRESS_NEST_HOST,
      port: Number(process.env.POSTGRES_NEST_PORT),
      username: process.env.POSTGRESS_NEST_USER,
      password: process.env.POSTGRES_NEST_PASSWORD,
      database: process.env.POSTGRES_NEST_DB,
      models: [CardsModel],
      autoLoadModels: true,
      sync: { alter: true },
    }),
    CardsModule,
  ],
  controllers: [],
  providers: [],
})
export class AppModule {}
