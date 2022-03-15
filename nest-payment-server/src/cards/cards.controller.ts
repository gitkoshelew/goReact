import { Body, Controller, Get, Post } from '@nestjs/common';
import { CreateCardDto } from './dto/create-card.dto';
import { CardsService } from './cards.service';
import { ApiOperation, ApiResponse, ApiTags } from '@nestjs/swagger';
import { CardsModel } from './cards.model';

@ApiTags('Cards')
@Controller('api/cards')
export class CardsController {
  constructor(private cardsService: CardsService) {}
  @ApiOperation({ summary: 'Create card' })
  @ApiResponse({ status: 200, type: CardsModel })
  @Post()
  create(@Body() cardDto: CreateCardDto) {
    return this.cardsService.createCard(cardDto);
  }
  @ApiOperation({ summary: 'Get all cards' })
  @ApiResponse({ status: 200, type: [CardsModel] })
  @Get()
  getAll() {
    return this.cardsService.getAllCards();
  }
}
