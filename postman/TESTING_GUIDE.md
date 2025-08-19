# Master Data Location API - Testing Guide

## 🎯 **Overview**

This guide provides comprehensive instructions for testing the Master Data Location API endpoints using the provided POSTMAN collection.

## 📋 **Issues Fixed**

### **Issue 1: Database Error on Location List Endpoint**
**Problem**: Route conflict where `/master/locations/stats` was being matched as `/master/locations/:id`
**Solution**: Reordered routes to place static routes before parameterized routes

### **Issue 2: Validation Error on Stats Endpoint**  
**Problem**: Stats endpoint was being treated as a location ID parameter
**Solution**: Moved stats endpoint before the `:id` route in the route definition

## 🚀 **Setup Instructions**

### **1. Import POSTMAN Collection**
1. Open POSTMAN
2. Click "Import" button
3. Select the `Master_Data_Location_API.postman_collection.json` file
4. The collection will be imported with all endpoints and examples

### **2. Configure Environment Variables**
The collection includes these variables that you can customize:
- `baseUrl`: API base URL (default: http://localhost:3000)
- `jwt_token`: JWT authentication token (will be set automatically after login)
- `province_id`, `regency_id`, `district_id`, `village_id`: Auto-populated during testing

### **3. Authentication Setup**
1. Run the "Login to Get JWT Token" request first
2. Update the email and password in the request body as needed
3. The JWT token will be automatically saved to the collection variables
4. All subsequent requests will use this token for authentication

## 📊 **Testing Workflow**

### **Recommended Testing Order:**

#### **Phase 1: Authentication & Statistics**
1. **Login to Get JWT Token** - Get authentication token
2. **Get Location Statistics** - Verify stats endpoint works

#### **Phase 2: Create Location Hierarchy**
3. **Create Province** - Creates a province and saves ID
4. **Create Regency** - Creates a regency under the province
5. **Create District** - Creates a district under the regency  
6. **Create Village** - Creates a village under the district

#### **Phase 3: Read Operations**
7. **Get All Locations (Mixed Types)** - Test listing without type filter
8. **Get All Provinces** - Test province listing with pagination
9. **Get Province by ID** - Test single province retrieval
10. **Get Regencies by Province** - Test hierarchical filtering
11. **Get District by ID** - Test hierarchical data retrieval
12. **Search All Locations** - Test search functionality

#### **Phase 4: Update Operations**
13. **Update Province** - Test partial updates
14. **Update Regency** - Test location updates

#### **Phase 5: Dependency & Deletion**
15. **Check Village Dependencies** - Test dependency checking
16. **Delete Village** - Test deletion (should work - no dependencies)
17. **Delete District** - Test deletion (should work after village deleted)
18. **Delete Regency** - Test deletion (should work after district deleted)
19. **Delete Province** - Test deletion (should work after regency deleted)

## 🔍 **Expected Response Formats**

### **Success Response Structure:**
```json
{
    "status": "success",
    "statusCode": 200,
    "payload": {
        // Response data here
    }
}
```

### **Error Response Structure:**
```json
{
    "status": "error", 
    "statusCode": 400,
    "payload": {
        "code": "validation001",
        "en": "validation001",
        "id": "validation001"
    }
}
```

### **Location Object Structure:**
```json
{
    "id": "uuid-string",
    "name": "Location Name",
    "code": "LOCATION_CODE", 
    "type": "province|regency|district|village",
    "parent_id": "parent-uuid-or-null",
    "latitude": -3.6954,
    "longitude": 128.1814,
    "regency_type": "kabupaten|kota", // Only for regencies
    "village_type": "desa|kelurahan",  // Only for villages
    "is_active": true,
    "created_at": "2025-02-27T10:30:00Z",
    "updated_at": "2025-02-27T11:00:00Z",
    "parent": {
        // Hierarchical parent information
    }
}
```

### **Paginated List Response:**
```json
{
    "data": [
        // Array of location objects
    ],
    "pagination": {
        "page": 1,
        "limit": 10,
        "total": 150,
        "total_pages": 15
    }
}
```

## 🧪 **Test Cases by Endpoint**

### **1. POST /master/locations (Create Location)**

#### **Valid Test Cases:**
- Create province (no parent_id required)
- Create regency (requires parent_id and regency_type)
- Create district (requires parent_id)
- Create village (requires parent_id and village_type)

#### **Invalid Test Cases:**
- Missing required fields (name, code, type)
- Invalid location type
- Missing parent_id for child locations
- Missing regency_type for regencies
- Missing village_type for villages
- Invalid coordinate values (lat/lng out of range)

### **2. GET /master/locations (List Locations)**

#### **Query Parameter Tests:**
- No parameters (returns all location types)
- `type=province` (returns only provinces)
- `type=regency&parent_id=uuid` (regencies in specific province)
- `search=Ambon` (search by name/code)
- `page=2&limit=5` (pagination)
- `sort_by=name&sort_order=desc` (sorting)

### **3. GET /master/locations/:id (Get Single Location)**

#### **Test Cases:**
- Valid ID with correct type parameter
- Invalid ID (should return 404)
- Valid ID with wrong type parameter
- Missing type parameter

### **4. PUT /master/locations/:id (Update Location)**

#### **Test Cases:**
- Update single field (partial update)
- Update multiple fields
- Update with invalid data
- Update non-existent location

### **5. DELETE /master/locations/:id (Delete Location)**

#### **Test Cases:**
- Delete location without dependencies (should succeed)
- Delete location with child locations (should fail)
- Delete location with associated reports (should fail)
- Delete non-existent location

## 🔧 **Troubleshooting**

### **Common Issues:**

#### **Authentication Errors:**
- Ensure JWT token is valid and not expired
- Check that Cookie header is properly set
- Verify login credentials are correct

#### **Validation Errors:**
- Check required fields are provided
- Verify data types and formats
- Ensure hierarchical relationships are correct

#### **Database Errors:**
- Verify database connection is working
- Check that referenced parent IDs exist
- Ensure foreign key constraints are satisfied

#### **Route Not Found:**
- Verify the correct HTTP method is used
- Check URL path is correct
- Ensure routes are properly registered

## 📈 **Performance Testing**

### **Load Testing Scenarios:**
1. **Bulk Creation**: Create 100+ locations in sequence
2. **Concurrent Reads**: Multiple simultaneous list requests
3. **Deep Hierarchy**: Test with 4-level location hierarchies
4. **Large Datasets**: Test pagination with 1000+ records

### **Expected Performance:**
- Single location operations: < 100ms
- List operations (10 items): < 200ms
- Search operations: < 300ms
- Bulk operations: < 1s per 10 items

## 🎯 **Success Criteria**

### **All Tests Should:**
- Return appropriate HTTP status codes
- Follow consistent response format
- Handle errors gracefully
- Maintain data integrity
- Respect hierarchical relationships
- Provide proper validation messages

### **Specific Validations:**
- ✅ Authentication required for all endpoints
- ✅ Hierarchical data properly nested
- ✅ Pagination working correctly
- ✅ Search functionality operational
- ✅ Dependency checking prevents invalid deletions
- ✅ Partial updates work correctly
- ✅ Error messages are descriptive

## 📝 **Notes**

- Always test in the recommended order to maintain data consistency
- Use the dependency checking endpoint before attempting deletions
- The collection automatically manages IDs between related requests
- All coordinates should be valid latitude/longitude values
- Location codes should follow Indonesian administrative coding standards
