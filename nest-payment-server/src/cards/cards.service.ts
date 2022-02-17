import { Injectable } from '@nestjs/common';

import { CardsModel } from './cards.model';
import { InjectModel } from '@nestjs/sequelize';
import { CreateCardDto } from './dto/create-card.dto';
import { encodeCvv } from '../utils/bcrypt';

@Injectable()
export class CardsService {
  constructor(
    @InjectModel(CardsModel) private cardRepository: typeof CardsModel,
  ) {}

  async createCard(dto: CreateCardDto) {
    const cvv = encodeCvv(dto.cvv);
    const card = await this.cardRepository.create({ ...dto, cvv });
    return card;
  }

  async getAllCards() {
    const cards = await this.cardRepository.findAll();
    return cards;
  }
}
