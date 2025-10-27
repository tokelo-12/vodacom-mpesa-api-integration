![M-pesa Logo](https://pbs.twimg.com/media/DYqHiduXUAAygx3.jpg:large)

## 🐘 Vodacom M-Pesa API Integration Demo
This repository provides a comprehensive demonstration of integrating with the Vodacom M-Pesa API, covering the crucial security and transaction stages required to successfully execute payments.

## ✨ Key Features
This project focuses on the core steps necessary for a successful M-Pesa integration:

     🔒 1. Application Key Encryption
        The initial step in securing your transactions is encrypting your application key.
        
        - Encryption Stage: Demonstrates the logic for encrypting your Application Key using the publicly available Base64 public key string provided on the Vodacom M-Pesa developer portal.
        
        - Security Focus: This crucial step ensures that your sensitive Application Key is protected before being sent to the M-Pesa service.

      🔑 2. Get Session Key Implementation
        Once the Application Key is secured, you can request a valid session key to authorize subsequent transactions.

        - Session Key Retrieval: Includes the full implementation of the getSessionKey request.

        - Authentication: Shows how to use your encrypted Application Key to obtain a temporary Session Key, which is required for all transaction requests.

      💰 3. C2B (Customer to Business) Transaction Demo
        The heart of the payment system—demonstrating how to initiate a transaction.
        
        - C2B Transaction: Contains the full implementation of a Customer to Business (C2B) transaction request.
        
        - Transaction Flow: Clearly shows how to use the retrieved Session Key to authorize and execute a mobile money payment.
        
