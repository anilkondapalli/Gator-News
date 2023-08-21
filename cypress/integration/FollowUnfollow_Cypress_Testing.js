describe("renders the home page",()=>{
    it("renders correctly",()=>{
        cy.visit("http://localhost:3001");
    cy.get('.content > :nth-child(3)').click();
    cy.get(':nth-child(1) > .form-control').type('alekhyagollam68@gmail.com');
    cy.get(':nth-child(2) > .form-control').type('123456');
    cy.get('form > .btn').click();
    cy.get('[href="/users"]').click();
    cy.get('.users__options > :nth-child(2)').click();
    cy.get('[href="/61fdb06e54134291baeb8ed1"] > .media-body > h5').click();
    cy.get('.unfollow > span').click();
    cy.get('.options > .btn').click();
    });
    });