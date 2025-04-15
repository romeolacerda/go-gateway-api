import { Module } from '@nestjs/common';
import { FraudService } from './fraud/fraud.service';
import { FrequentHighValueSpecification } from './fraud/specifications/frequent-high-value.specification';
import { SuspiciousAccountSpecification } from './fraud/specifications/suspicious-account.specification';
import { UnusualAmountSpecification } from './fraud/specifications/unusual-amount.specification';
import { FraudAggregateSpecification } from './fraud/specifications/fraud-aggregate.specification';
import { InvoicesService } from './invoices.service';
import { InvoicesController } from './invoices.controller';
import { InvoicesConsumer } from './invoices.consumer';
import * as kafkaLib from '@confluentinc/kafka-javascript';
import { PublishProcessedInvoiceListener } from './events/publish-processed-invoice.listener';
@Module({
  providers: [
    FraudService,
    FrequentHighValueSpecification,
    SuspiciousAccountSpecification,
    UnusualAmountSpecification,
    FraudAggregateSpecification,
    {
      provide: 'FRAUD_SPECIFICATIONS',
      useFactory: (
        frequentHighValueSpec: FrequentHighValueSpecification,
        suspiciousAccountSpec: SuspiciousAccountSpecification,
        unusualAmountSpec: UnusualAmountSpecification,
      ) => {
        return [
          frequentHighValueSpec,
          suspiciousAccountSpec,
          unusualAmountSpec,
        ];
      },
      inject: [
        FrequentHighValueSpecification,
        SuspiciousAccountSpecification,
        UnusualAmountSpecification,
      ],
    },
    InvoicesService,
    {
      provide: kafkaLib.KafkaJS.Kafka,
      useValue: new kafkaLib.KafkaJS.Kafka({
        'bootstrap.servers': 'kafka:29092',
      }),
    },
    PublishProcessedInvoiceListener,
  ],
  controllers: [InvoicesController, InvoicesConsumer],
})
export class InvoicesModule {}
