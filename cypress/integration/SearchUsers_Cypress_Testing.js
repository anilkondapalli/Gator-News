describe("renders the home page",()=>{
    it("renders correctly",()=>{
        cy.visit("http://localhost:3001");
    cy.get('.content > :nth-child(3)').click();
    cy.get(':nth-child(1) > .form-control').type('alekhyagollam68@gmail.com');
    cy.get(':nth-child(2) > .form-control').type('123456');
    cy.get('form > .btn').click();
    cy.get('[href="/users"]').click();
    cy.get('input').type('tej');
});
});