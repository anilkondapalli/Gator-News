describe("renders the home page",()=>{
    it("renders correctly",()=>{
        cy.visit("/");
        cy.get('.content > :nth-child(2)').click();
        cy.get(':nth-child(1) > .row > :nth-child(1) > .form-control').type('Santosh');
        cy.get(':nth-child(1) > .row > :nth-child(2) > .form-control').type('Maturi');
        cy.get('form > :nth-child(2) > .form-control').type('santoshmaturi@gmail.com');
        cy.get(':nth-child(3) > .row > :nth-child(1) > .form-control').type('123456');
        cy.get(':nth-child(3) > .row > :nth-child(2) > .form-control').type('123456');
        cy.get('form > .btn').click();
        
    });
    });