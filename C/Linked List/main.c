// Node structure.
struct Node
{
  int data;
  struct Node* next;
};

// Linked List ADT.
struct LinkedList {
  struct Node* head;
  int size;
};

// Functions to operate on the linked list.
void initializeList(struct LinkedList* list);
void insertAtBeginning(struct LinkedList* list, int data);
void insertAtEnd(struct LinkedList* list, int data);
void insertAtPosition(struct LinkedList* list, int data, int position);
void deleteFromBeginning(struct LinkedList* list);
void deleteFromEnd(struct LinkedList* list);
void deleteFromPosition(struct LinkedList* list, int position);
void displayList(struct LinkedList* list);
