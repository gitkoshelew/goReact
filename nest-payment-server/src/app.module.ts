import { Module } from '@nestjs/common';
import { CardsModule } from './cards/cards.module';

@Module({
  imports: [CardsModule],
  controllers: [],
  providers: [],
})
export class AppModule {}
