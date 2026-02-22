# Industry Templates

**go-microx** offers pre-configured templates for common industry use cases. These templates are not just empty folders; they come with basic service implementations, Docker configurations, and a central API Gateway.

## E-commerce Template

Perfect for building the next big marketplace. It scaffolds the following services:

- **User Service**: Registration, authentication, and profile management.
- **Product Service**: Catalog management, inventory tracking.
- **Order Service**: Order placement, status tracking.
- **Payment Service**: Integration logic for payment providers.
- **Cart Service**: User shopping cart management.
- **API Gateway**: Central entry point for all frontend requests.

## Video Streaming Template

Designed for high-traffic media applications:

- **User Service**: Profile and subscription management.
- **Video Service**: Meta-data management and video catalog.
- **Transcoding Service**: Logic for converting videos for different devices.
- **Recommendation Service**: Personalized content feed logic.
- **API Gateway**: Handles request routing and rate limiting.

## Food Delivery Template

Complete logistics and ordering system:

- **Restaurant Service**: Management of menus and restaurant profiles.
- **Order Service**: Real-time order processing.
- **Delivery Service**: Courier assignment and tracking logic.
- **Payment Service**: Secure transaction handling.
- **API Gateway**: Unified API for customers and restaurants.

## Custom Template

If your project doesn't fit these categories, use the **Custom** template. It provides a clean, well-structured Go workspace where you can define your own microservices while still benefiting from **go-microx**'s configuration and deployment scaffolding.

---

### Selecting a Template

During the `go-microx new` flow, you will be prompted to select one of these. Simply use the arrow keys and press **Enter**!
