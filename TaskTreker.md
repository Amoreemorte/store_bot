# Current:

Tasks:

- [ ] Create card adding in moderate bot by command "\create_card"

    - [x] Create card handler

    - [x] Create card repository

    - [x] Add Isfinished and ModeratorId in Card struct

    - [ ] Add name(after command message)

        - [ ] Set moderator state "AddName"
    
        - [ ] Create message with texts "Напишите название товара(не более % символов): "


    - [ ] Select collection(after message)
        
        - [ ] Validate name
        
        - [ ] Set moderator state "SelectColection"
        
        - [ ] Get coolections and set them in context
        
        - [ ] Create card with name
        
        - [ ] Create message with texts "Выберите коллекцию: " and buttons by collections

    - [ ] If name is incorrect repeat with incorrect_name message


    - [ ] Add desсription(after callback addCollection< sep >collection_name)
        
        - [ ] Set moderator state "AddDesription"
        
        - [ ] Get card by moderatorId
        
        - [ ] Update card with collection
        
        - [ ] Create message with texts "Напишите описание товара(не более % символов): "

    - [ ] Add price(after message)
        
        - [ ] Validate desription
        
        - [ ] Set moderator state "AddPrice"
        
        - [ ] Get card by moderatorId
        
        - [ ] Update card with description
        
        - [ ] Create message with texts "Напишите цену товара(не более %): "
    
    - [ ] If description is incorrect repeat with incorrect_description message

    - [ ] Add photo(after message)
        
        - [ ] Validate Price
        
        - [ ] Set moderator state "AddPhoto"
        
        - [ ] Get card by moderatorId
        
        - [ ] Update card with price
        
        - [ ] Create message with texts "Отправьте фото одним сообщением(не более %, остальные будут проигнорированы): "
    
    - [ ] If price is incorrect repeat with incorrect_price message

    - [ ] End card creating(after message with photos)
        
        - [ ] Get card by moderatorId
        
        - [ ] Set card IsFinished = true
        
        - [ ] Reset moderator state(set "NoAction")
<!-- Think with complectation of card -->

# Ended: 
